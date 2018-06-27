package numbers;

import java.util.Scanner;

public class CalculateCircumference {
	public static void main(String[] args) {
		float radius = userInputRadius();
		double circumference = calculateCircumference(radius);
		System.out.println("The circumference is: " + circumference);
	}

	public static float userInputRadius() {
		Scanner reader = new Scanner(System.in);
		System.out.println("Please input the radius: ");
		float radius = reader.nextFloat();
		reader.close();
		return radius;
	}

	public static double calculateCircumference(float radius) {
		return radius * 2 * Math.PI;
	}
}
