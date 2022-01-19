-- ALU : Arithmetic and Logical Unit --

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity ALU is
  port(
  AIn, BIn : in std_logic_vector(31 downto 0);
  ALUPlus, ALUMinus : in std_logic;
  ALUOut : out std_logic_vector(31 downto 0)
);
end ALU;

architecture RTL of ALU is

begin
  ALUOut <= AIn + BIn when (ALUPlus = '1') else
            AIn - BIn when (ALUMinus = '1') else
            (others => '0');

end RTL;
