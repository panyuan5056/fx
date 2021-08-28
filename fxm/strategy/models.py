import requests
import json
from django.db import models
from online.models import Online

#策略表
class Strategy(models.Model):
    class1   = models.CharField("分类1", max_length=200)
    class2   = models.CharField("分类2", max_length=200)
    class3   = models.CharField("分类3", max_length=200)
    name     = models.CharField("名称", max_length=200)
    dept     = models.CharField("深度", max_length=200)
    category = models.CharField("解析类型", max_length=200, default=1, choices=(("1","正则"),("2","规则"),("3","在之中"),("4","等于")))
    plan     = models.TextField("解析方法")
    desc          = models.TextField("描述", blank=True, null=True)
    forward       = models.CharField("涉及字段",max_length=20)
    create_time   = models.DateTimeField(auto_now=True)  
    update_time   = models.DateTimeField(auto_now_add=True)  

    def __str__(self):
        return self.name

    def save(self, *args, **kwargs):
        for online in Online.objects.all():
            payload = {
                        "class1":self.class1,
                        'class2':self.class2,
                        "class3":self.class3,
                        "name":self.name,
                        "dept":self.dept,
                        "category":self.category,
                        "plan":self.plan,
                        "desc":self.desc,
                        "forward":self.forward
                    }
            body = requests.post(online.server + '/api/v1/strategy/add', headers={"token":online.token}, data=json.dumps(payload))
            print(body.status_code)
        super(Strategy, self).save(*args, **kwargs)

    class Meta:
        verbose_name = '分级分类'
        verbose_name_plural = '分级分类'

    
