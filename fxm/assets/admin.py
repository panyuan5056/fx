try:
  import simplejson as json
except Exception as e:
  import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import Assets
#from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
 

class AssetsAdminForm(forms.ModelForm):
  class Meta:
      model = Assets
      widgets = {} 
      fields = '__all__'
  
  def __init__(self, *args, **kwargs):
      super(AssetsAdminForm, self).__init__(*args, **kwargs)
  
class AssetsAdmin(admin.ModelAdmin):
    form = AssetsAdminForm 
    list_display = ('name', 'ip', 'category', 'zhongxin', 'bumen', 'xitong', 'status', 'update_time')
    search_fields = ('name',)
    def get_form(self, request, obj=None, **kwargs):
       form = super(AssetsAdmin, self).get_form(request, obj, **kwargs)
       return form

    def get_queryset(self, request):
        qs = super(AssetsAdmin, self).get_queryset(request).all()
        return qs


admin.site.register(Assets, AssetsAdmin)

