defmodule Main do
  def run do
    a1 = Dog.start_link("Rex")
    a2 = Cat.start_link("Whiskers")

    send(a1, {:speak, self()})
    send(a2, {:speak, self()})

    receive do
      {:spoken, msg} -> IO.puts(msg)
    end

    receive do
      {:spoken, msg} -> IO.puts(msg)
    end

    send(a1, :stop)
    send(a2, :stop)
  end
end
