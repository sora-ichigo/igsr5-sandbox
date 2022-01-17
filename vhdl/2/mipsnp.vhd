-- Top Module of MIPS Non Pipeline --完成版

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity MIPSnp is
    port(
        CLK, RESET : in std_logic;
        PCOut : out std_logic_vector(7 downto 0)
    );
end MIPSnp;

architecture RTL of MIPSnp is
component IM is
    port(
        CLK : in std_logic;
        IMA : in std_logic_vector(31 downto 0);
        IMOut : out std_logic_vector(31 downto 0)
    );
end component;

component MP is
    port(
        CLK : in std_logic;
        OP : in std_logic_vector(5 downto 0);
        MPC : out std_logic_vector(15 downto 0)
    );
end component;

component RF is
    port(
        CLK, RESET : in std_logic;
        Rstd : in std_logic_vector(14 downto 0);  -- rs, rt, rd
        RSel : in std_logic_vector(1 downto 0);  -- rsSel, rtSel
        rC : in std_logic_vector(31 downto 0);
        RSet : in std_logic_vector(2 downto 0);  -- rdSet, rtSet, r31Set
        rA, rB : out std_logic_vector(31 downto 0)
    );
end component;

component ALU is
    port(
        AIn, BIn : in std_logic_vector(31 downto 0);
        ALUPlus, ALUMinus : in std_logic;
        ALUOut : out std_logic_vector(31 downto 0)
    );
end component;

component DM is
    port(
        CLK : in std_logic;
        DMA : in std_logic_vector(31 downto 0); 
        DMIn : in std_logic_vector(31 downto 0);
        DMOut : out std_logic_vector(31 downto 0);
        MRead, MWrite : in std_logic
    );
end component;

signal pc4 : std_logic_vector(31 downto 0);

-- IF
signal pc : std_logic_vector(31 downto 0);  -- Address
signal IMOut : std_logic_vector(31 downto 0);  -- Instruction
signal op : std_logic_vector(5 downto 0);
signal rsrtrd : std_logic_vector(14 downto 0);
signal imm : std_logic_vector(31 downto 0);
signal addr : std_logic_vector(25 downto 0);

-- MP
signal MPC : std_logic_vector(15 downto 0);  -- MP Control
signal pc4addrSel, addrSel, rASel : std_logic;  -- IF
signal rsSel, rtSel : std_logic;  -- ID
signal ALUPlus, ALUMinus, iaSel : std_logic;  -- EX
signal MRead, MWrite : std_logic;  -- MEM
signal rdSet, rtSet, r31Set, EXSel, MEMSel, pc4Sel : std_logic;  -- WB

-- ID
signal Rstd : std_logic_vector(14 downto 0);  -- rs, rt, rd
signal RSel : std_logic_vector(1 downto 0);  -- rsSel, rtSel
signal RSet : std_logic_vector(2 downto 0);  -- rdSet, rtSet, r31Set
signal rA, rB, rC : std_logic_vector(31 downto 0);

-- EX
signal AIn, BIn, ALUOut : std_logic_vector(31 downto 0);
signal ALUzero : std_logic;

-- MEM
signal DMA : std_logic_vector(31 downto 0);
signal DMIn, DMOut : std_logic_vector(31 downto 0);

begin

    PCout <= pc(7 downto 0);
    pc4 <= pc + 4;

-- IF : Instruction Fetch
    process(CLK, RESET)
    variable sel : std_logic_vector(2 downto 0);
    begin
        if (RESET = '0') then
            pc <= (others => '0');
        elsif (CLK'event and CLK = '1') then
            sel := (pc4addrSel and ALUZero) & addrSel & rASel;
            case (sel) is
				when "100" => pc <= pc4 +(imm(29 downto 0) & "00");
				when "010" => pc <= "0000" & addr & "00";
				when "001" => pc <= rA;
                when others => pc <= pc4;
            end case;
        end if;
    end process;

    IMi : IM port map (CLK, pc, IMOut);  -- IM Instance
    op <= IMOut(31 downto 26);
    rsrtrd <= IMOut(25 downto 11);
    imm <= X"FFFF" & IMOut(15 downto 0) when (IMOut(15) = '1') else  -- imm
           X"0000" & IMOut(15 downto 0);  -- 32 bit Signal
    addr <= IMOut(25 downto 0);  -- 26 bit

-- MP : Micro Program
    MPi : MP port map (CLK, op, MPC);  -- MP Instance
    -- IF
    pc4addrSel <= MPC(15);  -- pc + 4 + addr Select
    addrSel    <= MPC(14);  -- addr Select
    rASel      <= MPC(13);  -- rA Select
    -- ID
    rsSel      <= MPC(12);  -- rs Select
    rtSel      <= MPC(11);  -- rt Select
    -- EX
    ALUPlus    <= MPC(10);  -- ALU +
    ALUMinus   <= MPC(9);   -- ALU -
    iaSel      <= MPC(8);   -- imm/addr Select
    -- MEM
    MRead      <= MPC(7);   -- Memory Read
    MWrite     <= MPC(6);   -- Memory Write
    -- WB
    EXSel      <= MPC(5);   -- EX Select
    MEMSel     <= MPC(4);   -- MEM Select
    pc4Sel     <= MPC(3);   -- pc + 4 Select
    rdSet      <= MPC(2);   -- rd Set
    rtSet      <= MPC(1);   -- rt Set
    r31Set     <= MPC(0);   -- r31 Set

-- ID : Instruction Decode
    Rstd <= rsrtrd;
    RSel <= rsSel & rtSel;
    RSet <= rdSet & rtSet & r31Set;
    RFi : RF port map (CLK, RESET, Rstd, Rsel, rC, RSet, rA, rB);  -- RF Instance

-- EX : Execution
    AIn <= rA;  -- rA
    BIn <= imm when (iaSel = '1') else  -- Immediate
           rB;  -- rB
    ALUi : ALU port map (AIn, BIn, ALUPlus, ALUMinus, ALUOut);  -- ALU Instance
    ALUZero <= '1' when (ALUOut = 0) else
               '0';

-- MEM : Memory
    DMA <= ALUOut;
    DMIn <= rB;
    DMi : DM port map (CLK, DMA, DMIn, DMOut, MRead, MWrite);  -- DM Instance

-- WB : Write Back
    rC <= ALUOut when (EXSel = '1') else
    	  DMOut when (MEMSel = '1') else
    	  pc4 when (pc4Sel = '1') else
    	  (others => '0');
    	  
end RTL;