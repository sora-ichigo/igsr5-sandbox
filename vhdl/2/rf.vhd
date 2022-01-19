-- RF : Register File --

library IEEE;
use IEEE.std_logic_1164.all;
use IEEE.std_logic_unsigned.all;

entity RF is
  port(
  CLK, RESET : in std_logic;
  Rstd : in std_logic_vector(14 downto 0);  -- rs, rt, rd
  RSel : in std_logic_vector(1 downto 0);  -- rsSel, rtSel
  rC : in std_logic_vector(31 downto 0);
  RSet : in std_logic_vector(2 downto 0);  -- rdSet, rtSet, r31Set
  rA, rB : out std_logic_vector(31 downto 0)
);
end RF;

architecture RTL of RF is
  signal rs, rt, rd : std_logic_vector(4 downto 0);
  signal rsSel, rtSel : std_logic;
  signal r1, r2, r3, r29, r30, r31 : std_logic_vector(31 downto 0);
  signal rdSet, rtSet, r31Set : std_logic;

begin

  rs <= Rstd(14 downto 10);
  rt <= Rstd(9 downto 5);
  rd <= Rstd(4 downto 0);
  rsSel <= RSel(1);
  rtSel <= RSel(0);
  rdSet <= RSet(2);
  rtSet <= RSet(1);
  r31Set <= RSet(0);

  RFi : process(CLK, RESET)
  begin
    if (RESET = '0') then
      r1 <= X"11111111";  -- r1 for test
      r2 <= X"22222222";  -- r2 for test
      r3 <= X"33333333";  -- r3 for test
      r29 <= (others => '0');  -- sp
      r30 <= (others => '0');  -- fp
      r31 <= (others => '0');  -- ra
    elsif (CLK'event and CLK = '1') then
      if (rdSet = '1') then
        -- NOTE: この分岐は内部実装
        -- 今回の MIPS ではレジスタファイルのうち
        -- r1, r2, r3, r29, r30, r31 のアドレスいずれかが使用される
        -- その前提で条件分岐をハードコーディングしている
        case (rd) is
          when "00001" => r1 <= rC;
          when "00010" => r2 <= rC;
          when "00011" => r3 <= rC;
          when "00100" => r29 <= rC;
          when "00101" => r30 <= rC;
          when "00110" => r31 <= rC;
          when others => NULL;
        end case;
      elsif (rtSet = '1') then
        case (rt) is
          when "00001" => r1 <= rC;
          when "00010" => r2 <= rC;
          when "00011" => r3 <= rC;
          when "00100" => r29 <= rC;
          when "00101" => r30 <= rC;
          when "00110" => r31 <= rC;                
          when others => NULL;
        end case;
      elsif (r31Set = '1') then
        r31 <= rC;
      end if;
    end if;
  end process;

  RAo : process(rsSel, rs, r1, r2, r3, r29, r30, r31)
  begin
    if (rsSel = '1') then  --rA
      case (rs) is
        when "00001" => rA <= r1;
        when "00010" => rA <= r2;
        when "00011" => rA <= r3;
        when "00100" => rA <= r29;
        when "00101" => rA <= r30;
        when "00110" => rA <= r31;  
        when others => rA <= (others => '0');  -- r0
      end case;
    else
      rA <= (others => '0');
    end if;
  end process;

  RBo : process(rtSel, rt, r1, r2, r3, r29, r30, r31)
  begin
    if (rtSel = '1') then  -- rB
      case (rt) is
        when "00001" => rB <= r1;
        when "00010" => rB <= r2;
        when "00011" => rB <= r3;
        when "00100" => rB <= r29;
        when "00101" => rB <= r30;                
        when "00110" => rB <= r31;
        when others => rB <= (others => '0');  -- r0
      end case;
    else
      rB <= (others => '0');
    end if;
  end process;
end RTL;
