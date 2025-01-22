import at.favre.lib.crypto.bcrypt.BCrypt;
import org.apache.commons.lang3.RandomStringUtils;

public class BcryptAtFavre {

    public static void main(String[] args) {
        // 18 + 1 + 55 = 74, so above 72 characters' limit of BCrypt
        var userId = RandomStringUtils.randomAlphanumeric(18);
        var username = RandomStringUtils.randomAlphanumeric(55);
        var password = "super-duper-secure-password";

        var combinedString = String.format("%s:%s:%s", userId, username, password);

        var combinedHash = BCrypt.withDefaults().hashToString(12, combinedString.toCharArray());

        // let's try to break it
        var wrongPassword = "wrong-password";
        var wrongCombinedString = String.format("%s:%s:%s", userId, username, wrongPassword);

        var result = BCrypt.verifyer().verify(combinedHash.toCharArray(), wrongCombinedString);
        if (result.verified) {
            System.out.println("Password is correct");
        } else {
            System.out.println("Password is incorrect");
        }
    }
}
