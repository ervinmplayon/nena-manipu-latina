# Tempoorary Tattoos: Will Get Moved Eventually
* In Go, interfaces should live near where they're used, not abstracted away into a generic directory
* Go idiom: "Design with interfaces, but don't expoprt them unless you have to."
* `interfaces/` make sense when I am defining contracts across multiple subsystems and want to centralize them (e.g., in a large monorepo)
* `interfaces/` would make sense if I am building a plugin system where interface discovery is the goal. 