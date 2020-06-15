import datetime

from django.db import models
from django.utils import timezone


# Create your models here.
class Message(models.Model):
    private_key = models.CharField(max_length=16)
    message_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField('date published')
    def __str__(self):
        return self.message_text
