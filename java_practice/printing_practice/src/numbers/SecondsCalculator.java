package numbers;

import java.util.Scanner;

public class SecondsCalculator {
	public static void main(String[] args) {
		int days = userInputDays();
		int seconds = calculateSeconds(days);
		System.out.println("There are " + seconds + " seconds in " + days + " days!");
	}

	public static int calculateSeconds(int days) {
		int hours = days * 24;
		int minutes = hours * 60;
		int seconds = minutes * 60;
		return seconds;
	}

	public static int userInputDays() {
		Scanner reader = new Scanner(System.in);
		System.out.print("Please input days: ");
		int days = reader.nextInt();
		reader.close();
		return days;
	}
}
