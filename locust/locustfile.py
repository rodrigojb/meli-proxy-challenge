from locust import FastHttpUser, task

class HelloMeliUser(FastHttpUser):
    @task
    def hello_meli(self):
        self.client.get("/categories/MLA3531")
        self.client.get("/categories/MLA3530")
        self.client.get("/currencies")
        self.client.get("/classified_locations/countries")