class Main
  def dfs(i, sum)
    return sum == @k if i == @n
    return true if dfs(i+1, sum)
    return true if dfs(i+1, sum + @a[i])
    return false
  end

  def solve()
    @t = gets.chomp.to_i
    @t.times do
      @n = gets.chomp.to_i
      @a = gets.chomp.split(' ').map(&:to_i)
      @k = gets.chomp.to_i
      if (self.dfs(0, 0)) then
        puts 'Yes'
      else
        puts 'No'
      end
    end
  end
end

if $0 == __FILE__
  x = Main.new()
  x.solve()
end