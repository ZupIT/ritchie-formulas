package ${package_name}

import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get

@Controller
class MyController(private val myService: MyService) {
    @Get("/action")
    fun getAction() = myService.createAction()

    @Get("/screen")
    fun getScreen() = myService.createScreen()

    @Get("/builder")
    fun getScreenBuilder() = myService.createScreenBuilder()

    @Get("/widget")
    fun getWidget() = myService.createWidget()
}