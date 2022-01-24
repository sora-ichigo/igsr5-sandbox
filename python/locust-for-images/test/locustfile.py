from locust import HttpUser, task

class GetImage(HttpUser):
    @task
    def get_image(self):
        # self.client.get("/PwwkcCA")
        # self.client.get("/qwksUBo?width=600&height=350&format=png")
        # self.client.get("/h7d2yZ3?width=1136&format=webp")
        # self.client.get("/gMJFFpz?width=600&height=350&format=png")
        # self.client.get("/ping")

        # 274 KB
        # self.client.get("/rkB8ePL?width=600&height=300")
        # 1,200 KB
        # self.client.get("/LXPdsTx?width=900&height=500&format=png")
        # 3,359 KB
        self.client.get("/aKnmkdN?width=900&height=500&format=png")
