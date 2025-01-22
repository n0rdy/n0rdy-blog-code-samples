import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.security.crypto.bcrypt.BCrypt;

public class BcriptSpringSecurity {
    public static void main(String[] args) {
        // 18 + 55 + 1 = 74, so above 72 characters' limit of BCrypt
        var userId = RandomStringUtils.randomAlphanumeric(18);
        var username = RandomStringUtils.randomAlphanumeric(55);
        var password = "super-duper-secure-password";

        var combinedString = String.format("%s:%s:%s", userId, username, password);

        var combinedHash = BCrypt.hashpw(combinedString, BCrypt.gensalt());

        // let's try to break it
        var wrongPassword = "wrong-password";
        var wrongCombinedString = String.format("%s:%s:%s", userId, username, wrongPassword);

        if (BCrypt.checkpw(wrongCombinedString, combinedHash)) {
            System.out.println("Password is correct");
        } else {
            System.out.println("Password is incorrect");
        }
    }
}

