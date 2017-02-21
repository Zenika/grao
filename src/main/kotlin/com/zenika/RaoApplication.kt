package com.zenika


import com.fasterxml.jackson.module.kotlin.KotlinModule
import org.slf4j.LoggerFactory
import org.springframework.boot.SpringApplication
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.context.annotation.Bean
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder
import springfox.documentation.builders.ApiInfoBuilder
import springfox.documentation.builders.PathSelectors
import springfox.documentation.builders.RequestHandlerSelectors
import springfox.documentation.service.ApiInfo
import springfox.documentation.spi.DocumentationType
import springfox.documentation.spring.web.plugins.Docket
import springfox.documentation.swagger2.annotations.EnableSwagger2

@SpringBootApplication
@EnableSwagger2
open class RaoApplication {
    private val log = LoggerFactory.getLogger(RaoApplication::class.java)

    @Bean
    open fun objectMapperBuilder()
            = Jackson2ObjectMapperBuilder().modulesToInstall(KotlinModule())

    @Bean
    open fun api(): Docket {
        return Docket(DocumentationType.SWAGGER_2)
                .select()
                .apis(RequestHandlerSelectors.any())
                .paths(PathSelectors.any())
                .build()
                .apiInfo(apiInfo())
    }

    private fun apiInfo(): ApiInfo {
        return ApiInfoBuilder()
                .title("RAO API")
                .description("API for RAO application")
                .license("MIT License")
                .licenseUrl("https://github.com/Zenika/rao/blob/master/LICENSE")
                .version("1.0")
                .build()
    }




}

fun main(args: Array<String>) {
    SpringApplication.run(RaoApplication::class.java, *args)
}
