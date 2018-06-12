require "pry"
class Change

  def self.generate(coins, amount, position = 0, change = Array.new)
    return change.flatten.reverse if amount <= 0
    coins = coins.reverse
    divisor = amount / coins[position]
    amount %= coins[position]
    change << Array.new(divisor){coins[position]}
    position += 1
    generate(coins.reverse, amount, position, change)
  end

end
