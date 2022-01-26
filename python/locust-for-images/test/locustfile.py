from locust import HttpUser, task

class GetImage(HttpUser):
    @task
    def get_image(self):
        # self.client.get("/LXPdsTx?width=900&height=500&format=png")
        # self.client.get("/LXPdsTx?width=900&height=500")
        self.client.get("/LXPdsTx?format=png")
