from locust import HttpUser, task

class GetImage(HttpUser):
    @task
    def get_image(self):
        # self.client.get("/PwwkcCA")
        self.client.get("/qwksUBo?width=600&height=350&format=png")
        # self.client.get("/ping")
