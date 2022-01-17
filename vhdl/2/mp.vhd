-- MP : Micro Program --

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity MP is
    port(
        CLK : in std_logic;
        OP : in std_logic_vector(5 downto 0);
        MPC : out std_logic_vector(15 downto 0)
    );
end MP;

architecture RTL of MP is

begin
         -- IF+ID+EX+MEM+WB
         -- pc+4+addr
         -- |addr   Read
         -- ||rA    |Write
         -- |||rs   ||EX
         -- ||||rt  |||MEM
         -- |||||+-i||||pc+4
         -- |||||||||||||rd
         -- ||||||||||||||rt
         -- |||||||||||||||r31
         -- ||||||||||||||||
    MPC <= "0001110000100100" when (OP = "000000") else  -- add
           "0001010100100010" when (OP = "000001") else  -- addi
           "0001010110010010" when (OP = "000010") else  -- lw
           "0001110101000000" when (OP = "000011") else  -- sw
           "1001101000000000" when (OP = "000100") else  -- beq
           "0100000000000000" when (OP = "000101") else  -- j
           "0100000000001001" when (OP = "000110") else  -- jal
           "0011000000000000" when (OP = "000111") else  -- jr
           "0000000000000000";

end RTL;