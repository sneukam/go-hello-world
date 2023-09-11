/*
I'm learning about the Golang language.
Getting ready for the new job at Applied Insight.
I hope I don't get stuck in DevOps.
But the languages they want do speak to some actual backend.
I hope that the job description is not a disguise for a pure devops role.
Otherwise, I can work on personal projects and gain some nice skills there.

Go mod
Enabling dependency tracking.

When my code imports packages contained in other modules,
I manage those dependencies through my code's own module.
Note from me: It looks like this codebase is a 'module'
My module is defined by a go.mod file that tracks the modules that provide those packages.
go.mod stays with my code, including in the source repository.

To enable dependency tracking, use the command `go mod init <path/to/filename>`
This will create a go.mod file in the root directory the command was execute from,
and thus set up your repository to track packages in other modules, and manage your dependencies.

If my module needs to be published for others to use,
the module path is typically a repository location that my code is kept at (github).
Here, I made my module path github.com/sneukam/repo-name
This allows go tools to download my module.

I declared the main package here in this file.
A package is a way to group functions, and it's made up of all the files in the same directory.
The main function executes by default when running the main package.

Okay, I just downloaded the quote module.
This has a few basic functions as a tutorial for beginners.
Rather than manually entering the required module in the mod file,
I added it here and used the command `go mod tidy` in the terminal.
It identified the package here and updated the mod file.
This is the way to go because the mod file looks like it has some
somewhat complicated syntax. Go mod tidy is my friend.

*/

package main

import "fmt"
import "rsc.io/quote/v4"

func main() {
	fmt.Println("Hello world!")
	fmt.Println("Gottem'")
	fmt.Println("--")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
