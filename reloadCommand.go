package main

func reloadCommand(args []string) bool {
	wg.Add(1)
	go getBetaUser()
	wg.Wait()
	info("Reload done.")
	return true
}
