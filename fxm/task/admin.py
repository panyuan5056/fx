try:
  import simplejson as json
except Exception as e:
  import json
from datetime import datetime as dt
import requests
from django.contrib import admin
from django.conf import settings
from django import forms
from .models import Task, TaskDetail
#from flyadmin.widget.forms import SelectBoxWidget, TimelineWidget, EditorWidget, DateTimeWidget, UploadImagesWidget, InputNumberWidget, UploadFileWidget, StepsWidget, StepsNormalWidget
 
class TaskAdminForm(forms.ModelForm):
  class Meta:
      model = Task
      widgets = {
          "password":forms.PasswordInput
      } 
      fields = '__all__'

  def __init__(self, *args, **kwargs):
      super(TaskAdminForm, self).__init__(*args, **kwargs)
    
class TaskAdmin(admin.ModelAdmin):
    form = TaskAdminForm 
    list_display = ('name', 'category', 'username', 'network', 'status', 'server', 'port', 'database', 'charset', 'update_time')
    search_fields = ('name', )

    def get_form(self, request, obj=None, **kwargs):
       form = super(TaskAdmin, self).get_form(request, obj, **kwargs)
       return form

    def get_queryset(self, request):
        qs = super(TaskAdmin, self).get_queryset(request)
        return qs.filter()


admin.site.register(Task, TaskAdmin)



class TaskDetailAdminForm(forms.ModelForm):
  class Meta:
      model = TaskDetail
      widgets = {
        
      } 
      fields = '__all__'

  def __init__(self, *args, **kwargs):
      super(TaskDetailAdminForm, self).__init__(*args, **kwargs)
    
class TaskDetailAdmin(admin.ModelAdmin):
    form = TaskDetailAdminForm 
    list_display = ('task', 'table_name', 'name', 'config', 'class1', 'class2', 'class3', 'update_time')
    search_fields = ('table_name', )

    def get_form(self, request, obj=None, **kwargs):
       form = super(TaskDetailAdmin, self).get_form(request, obj, **kwargs)
       return form

    def get_queryset(self, request):
        qs = super(TaskDetailAdmin, self).get_queryset(request)
        return qs.filter()


admin.site.register(TaskDetail, TaskDetailAdmin)
 
