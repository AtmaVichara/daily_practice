package variables;

public class ChickenCount {
	public static void main(String[] args) {
		double pounds = 3;
		System.out.println("Chicken Breast in Pounds:");
		System.out.println(pounds);
		System.out.println("Chicken Breast in Kilograms:");
		System.out.println(convertPoundstoKg(pounds));
	}

	public static double convertPoundstoKg(double pounds) {
		return (pounds / 0.453592);
	}
}
