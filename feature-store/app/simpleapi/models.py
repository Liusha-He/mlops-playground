from django.db import models


class Item(models.Model):
    title = models.CharField(max_length=200)
    description = models.CharField(max_length=300)
    created_at = models.DateTimeField(auto_now_add=True)
