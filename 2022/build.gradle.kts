plugins {
    kotlin("jvm") version "1.7.22"
}

repositories {
    mavenCentral()
}
dependencies {
    implementation("junit:junit:4.13.1")
    implementation("org.junit.jupiter:junit-jupiter:5.8.1")
    implementation(kotlin("test"))
}

tasks {
    sourceSets {
        main {
            java.srcDirs("src")
        }
    }

    wrapper {
        gradleVersion = "7.6"
    }
}
