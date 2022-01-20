from locust import HttpUser, task

class GetImage(HttpUser):
    @task
    def get_image(self):
        # self.client.get("/PwwkcCA")
        # self.client.get("/qwksUBo?width=600&height=350&format=png")
        self.client.get("/h7d2yZ3?width=1136&format=webp")
        # self.client.get("/gMJFFpz?width=600&height=350&format=png")
        # self.client.get("/ping")
