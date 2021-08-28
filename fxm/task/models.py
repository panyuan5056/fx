import json
import requests
from django.db import models
from online.models import Online

#策略表
class Task(models.Model):
    name      = models.CharField("名称", max_length=200) 
    category  = models.CharField("类型", max_length=200) 
    username  = models.CharField("登录帐号", max_length=200, blank=True, null=True) 
    password  = models.CharField("登录密码", max_length=200, blank=True, null=True)
    network   = models.CharField("网络", default="TCP", max_length=200)
    server    = models.CharField("IP地址", max_length=200)
    port      = models.IntegerField("端口", blank=True, null=True)
    database  = models.CharField("数据库", max_length=200)
    charset   = models.CharField("编码", default="UTF8", max_length=200, blank=True, null=True)
    status    = models.IntegerField("状态", default=1, choices=((1, "未开始"),(2, "进行中"),(3, "完成"),(3, "异常")), editable=False)
    desc      = models.TextField("描述", blank=True, null=True)
    tid       = models.IntegerField('下发发现id', blank=True, null=True)
    version   = models.IntegerField("版本", default=1)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    def __str__(self):
        return self.name
    
    def save(self, *args, **kwargs):
        if not self.tid:
            for online in Online.objects.order_by("cpu"):
                payload = {
                            "category":self.category,
                            'username':self.username,
                            "password":str(self.password),
                            "network":self.network,
                            "server":self.server,
                            "port":str(self.port),
                            "database":self.database,
                            "charset":self.charset
                        }
                try:
                    body = requests.post(online.server + '/api/v1/db/find', headers={"token":online.token}, data=json.dumps(payload))
                    if body.status_code == 200:
                        self.status = 2
                        self.tid = body.json()['result']['id']
                    break
                except Exception as e:
                    print(e)
                    self.status = 3
        super(Task, self).save(*args, **kwargs)

    class Meta:
        verbose_name = '敏感数据发现'
        verbose_name_plural = '敏感数据发现'

#策略表详细表
class TaskDetail(models.Model):
    task          = models.ForeignKey('Task', to_field='id', on_delete=models.CASCADE, blank=True, null=True, verbose_name='敏感发现', related_name='Detail_task')
    name          = models.CharField("字段名", max_length=200) 
    table_name    = models.CharField("表名", max_length=200) 
    total         = models.CharField("总数", max_length=200) 
    step          = models.CharField("数量", max_length=200) 
    config        = models.CharField("类型名称", max_length=200) 
    class1        = models.CharField("一级分类", max_length=200) 
    class2        = models.CharField("二级分类", max_length=200) 
    class3        = models.CharField("三级分类", max_length=200) 
    version       = models.IntegerField("版本", default=1)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 
    
    class Meta:
        verbose_name = '敏感数据发现详细'
        verbose_name_plural = '敏感数据发现详细'



















