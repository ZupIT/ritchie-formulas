package ${package_name}

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class BffBeagleSpringApplication

fun main(args: Array<String>) {
	runApplication<BffBeagleSpringApplication>(*args)
}
