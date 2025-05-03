# Elixir Animal OOP Simulation

### What features in Elixir are used to implement objects and OO pillars?

Even though Elixir isn't object-oriented and doesn’t have classes or objects like Java, it's still possible to simulate some of the OO pillars using processes and message passing. Here’s how I implemented the main OO concepts:

- **Encapsulation**: Each animal is a separate Elixir process started using `spawn_link`. These processes each hold their own private state (the animal's name) and only expose it through specific message patterns like `{:get_name, caller}` or `{:speak, caller}`. This is kind of like private variables and public methods in a Java class.

- **Abstraction**: I created a base module called `Animal` that defines a basic `loop/1` function with a generic `:speak` behavior. Even though it's not used directly in this program, it shows how we could define an abstract interface in Elixir. The idea is that all animals respond to `:speak` and `:get_name`, so the message format acts like a shared interface.

- **Inheritance**: Elixir doesn’t have traditional inheritance, but I simulated it by copying and modifying the logic from `Animal` into the `Dog` and `Cat` modules. Each of those modules uses the same structure but overrides the `:speak` behavior with its own message. It’s like extending a class in Java and overriding a method.

- **Polymorphism**: I can send the same message (like `:speak`) to both a dog and a cat process, and they’ll each respond in their own way. This is just like how in Java, you can call `speak()` on an `Animal` variable and get different behavior depending on whether it’s a `Dog` or `Cat`.

### What parts of your Elixir code map to the Java example?

- The `Animal` module is like the abstract Java class `Animal` — it defines the interface and base behavior.
- The `Dog` and `Cat` modules are like subclasses that inherit from `Animal` and override the `speak` method.
- The `start_link` function in each module is like the constructor in Java that sets the animal’s name.
- The `loop/1` function in each process acts like a class that holds state and methods.
- The `Main.run/0` function is like the `main` method in Java — it creates the animals and tells them to speak.
