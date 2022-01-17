-- IM : Instruction Memory --

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity IM is
    port(
        CLK : in std_logic;
        IMA : in std_logic_vector(31 downto 0);
        IMOut : out std_logic_vector(31 downto 0)
    );
end IM;

architecture RTL of IM is

begin
                  -- R Type : < op ><rs ><rt ><rd ><sha><func>
                  -- I Type : < op ><rs ><rt ><   imm/addr   >
                  -- J Type : < op ><          addr          >
    IMo : process(IMA(7 downto 2))
    begin
        case (IMA(7 downto 2)) is
            when "000000" =>  -- op
                    IMOut <= "00000100000000010000000000000100";
            when "000001" =>  -- 
                    IMOut <= "00000000001000010001000000000000";
            when "000010" =>  -- 
                    IMOut <= "00001100010000011111111111111100";
            when "000011" =>  -- 
                    IMOut <= "00001000001000100000000000000000";
            when "000100" =>  -- 
                    IMOut <= "00010000001000100000000000000001";
            when "000101" =>  -- 
                    IMOut <= "00011111111000000000000000000000";
            when "000110" =>  -- 
                    IMOut <= "00011000000000000000000001010000";
            when "000111" =>  -- 
                    IMOut <= "00010100000000000000000000000000";
            when others =>  -- nop : add $0,$0,$0
                    IMOut <= "00000000000000000000000000000000";
        end case;
    end process;

end RTL;