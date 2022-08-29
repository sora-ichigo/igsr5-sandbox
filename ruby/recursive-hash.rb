a = {a: {b: :c}}

# ネストしたhashをフラットな配列に変換する
# 入力: {a: {b: :c}}
# 出力: [:a, :b, :c]
def any_to_flat_array(any_value, arr: [])
  unless any_value.is_a?(Hash)
    arr << any_value
    return arr
  end

  any_value.map do |k, v|
    arr << k
    if v.is_a?(Hash)
      any_to_flat_array(v, arr: arr)
    else
      arr << v
    end
  end

  arr
end

p any_to_flat_array(a)
p any_to_flat_array(:a)
