from locust import HttpUser, task

class GetImage(HttpUser):
    @task
    def get_image(self):
        # self.client.get("/PwwkcCA")
        self.client.get("/qwksUBo")
        self.client.get("/ping")
