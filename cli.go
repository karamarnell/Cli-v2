// Package cli provides a minimal example for creating and organizing command line

// cli application can be written as follows:
//   func main() {
//     cli.NewApp().Run(os.Args)
//   }
//
// Of course this application does not do much, so let's make this an actual application:
//   func main() {
//     app := cli.NewApp()
//     app.Name = "greet"
//     app.Usage = "say a greeting"
//     app.Action = func(c *cli.Context) error {
//       println("Greetings")
//       return nil
//     }
//
//     app.Run(os.Args)
//   }
package cli
