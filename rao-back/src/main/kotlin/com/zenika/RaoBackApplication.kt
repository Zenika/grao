package com.zenika

import com.fasterxml.jackson.module.kotlin.KotlinModule
import com.zenika.filestorage.DropboxConnector
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.CommandLineRunner
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
import javax.annotation.PostConstruct

@SpringBootApplication
@EnableSwagger2
open class RaoBackApplication {
    private val log = LoggerFactory.getLogger(RaoBackApplication::class.java)

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

    /**
     * TODO :
     * - P1: Lire le chemin racine des documents dans le fichier de config (ou variable d'env)
     * - P1: Au lancement : parcourir tous les documents et les indexer dans Algolia
     * - P1: Stocker en BDD les meta des fichiers déjà indexés (id, chemin relatif, nom)
     * - P1: Tous les soirs, parcourir les nouveaux documents et les indexer dans Algolia + maj la BDD
     *
     * - P2: services pour rechercher dans Algolia
     * - P2: offrir une IHM pour rechercher les documents via mots clés
     */

    @Bean
    open fun init(dropboxConnector:DropboxConnector): CommandLineRunner {
        return CommandLineRunner {
            dropboxConnector.connect()
            dropboxConnector.getNewFilesInformation()
        }
    }
}

fun main(args: Array<String>) {
    SpringApplication.run(RaoBackApplication::class.java, *args)
}
