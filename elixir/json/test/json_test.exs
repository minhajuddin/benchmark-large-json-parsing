defmodule JsonTest do
  use ExUnit.Case
  doctest Json

  test "poison" do
    {us, _} =
      :timer.tc(fn ->
        Path.relative_to("../../data/10mb.json", __ENV__.file)
        |> File.read!()
        |> Poison.decode!()
      end)

    IO.puts("Time taken [Poison]: #{us / 1000}ms")
  end

  test "jason" do
    {us, _} =
      :timer.tc(fn ->
        Path.relative_to("../../data/10mb.json", __ENV__.file)
        |> File.read!()
        |> Jason.decode!()
      end)

    IO.puts("Time taken [Jason]: #{us / 1000}ms")
  end
end
