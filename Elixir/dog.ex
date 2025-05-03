defmodule Dog do
  def start_link(name) do
    spawn_link(fn -> loop(name) end)
  end

  defp loop(name) do
    receive do
      {:get_name, caller} ->
        send(caller, {:name, name})
        loop(name)

      {:speak, _caller} ->
        IO.puts("#{name} says: Woof!")
        loop(name)

      :stop ->
        :ok

      _ ->
        loop(name)
    end
  end
end
