def is_permutation_palindrome(input)
  h = Hash.new

  input.each_char do |c|
    if h.key?(c)
      h[c] += 1
    else
      h[c] = 1
    end
  end

  odd_count = 0

  h.each_value do |v|
    if v.odd?
      odd_count += 1
    end

    return false if (h.size.even? && odd_count > 0) || (h.size.odd? && odd_count > 1)
  end

  true
end

result = is_permutation_palindrome("civil")

puts "result: #{result}"
