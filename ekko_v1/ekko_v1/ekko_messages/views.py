from django.shortcuts import render
from django.views.generic.list import ListView
from django.http import HttpResponse
from .models import Message

# Create your views here.

def index(request):
    message_list = Message.objects.order_by('-pub_date')[:5]
    output = ', '.join([m.message_text for m in message_list])
    return render(request, 'index.html')

def sucess(request):
    return render(request, 'success.html')
