package main.java.Main;

import common.java.JGrapeSystem.GscBooster;

public class Main {
	public static void main(String[] args) {
		System.out.println("<rd-user-service> service");
		GscBooster.start(args, () -> {
			 // load todo
		});
	}
}
