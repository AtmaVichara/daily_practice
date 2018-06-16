package printing_practice;

import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Spruce {
	public static void main(String[] args) {
		int d = 8;
		int ii = 1;
		for (int i = 0; i < 9; i++) {
			String whiteSpace = IntStream.range(0, d).mapToObj(x -> " ").collect(Collectors.joining(""));
			if (i == 0) {
				System.out.println(whiteSpace + "*" + whiteSpace);
			} else {
				String bangs = IntStream.range(0, (i + ii)).mapToObj(x -> "*").collect(Collectors.joining(""));
				System.out.println(whiteSpace + bangs + whiteSpace);
			}
			d--;
			ii++;
		}
	}
}
