-- DM : Data Memory --

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity DM is
    port(
        CLK : in std_logic;
        DMA : in std_logic_vector(31 downto 0); 
        DMIn : in std_logic_vector(31 downto 0);
        DMOut : out std_logic_vector(31 downto 0);
        MRead, MWrite : in std_logic
    );
end DM;

architecture RTL of DM is
type MEMORY is array (0 to 63) of std_logic_vector(31 downto 0);
signal DataMemory : MEMORY;

begin

    DMi : process(CLK)
    begin
        if (CLK'event and CLK = '1') then
            if (MWrite = '1') then  -- Memory Write
                DataMemory(conv_integer(DMA(7 downto 2))) <= DMIn;
            end if;
        end if;
    end process;
    -- Memory Read
    DMOut <= DataMemory(conv_integer(DMA(7 downto 2))) when (MRead = '1') else
             (others => '0');

end RTL;