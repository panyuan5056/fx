try:
  import simplejson as json
except Exception as e:
  import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import Strategy
#from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
 
class StrategyAdminForm(forms.ModelForm):
  class Meta:
      model = Strategy
      widgets = {} 
      fields = '__all__'

  def __init__(self, *args, **kwargs):
      super(StrategyAdminForm, self).__init__(*args, **kwargs)
    
class StrategyAdmin(admin.ModelAdmin):
    form = StrategyAdminForm 
    list_display = ('name', 'class1', 'class2', 'class3', 'dept', 'category', 'forward', 'update_time')
    search_fields = ('name', )

    def get_form(self, request, obj=None, **kwargs):
       form = super(StrategyAdmin, self).get_form(request, obj, **kwargs)
       return form

    def get_queryset(self, request):
        qs = super(StrategyAdmin, self).get_queryset(request)
        return qs.filter()


admin.site.register(Strategy, StrategyAdmin)
