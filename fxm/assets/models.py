from django.db import models
from flyadmin.views.charts import bar, pie, line

#策略表
class Assets(models.Model):
    name          = models.CharField("名称", max_length=200)
    ip            = models.GenericIPAddressField()
    category      = models.CharField("类型", max_length=200)
    xitong        = models.CharField("系统", max_length=200)
    bumen         = models.CharField("部门", max_length=200)
    zhongxin      = models.CharField("中心", max_length=200)
    status        = models.IntegerField("状态", choices=((1, "正常"), (0,"失效")), default=1)
    message       = models.TextField("描述", blank=True, null=True)
    create_time   = models.DateTimeField("创建时间",auto_now=True)  
    update_time   = models.DateTimeField("更新时间",auto_now_add=True) 

    def __str__(self):
        return self.name
        
    class Meta:
        verbose_name = '资产管理'
        verbose_name_plural = '资产管理'

    @classmethod
    def show_plots(cls):
        result = []
        assets = cls.objects.all()
        xx = {}
        for asset in assets:
            if xx.get(asset.category):
                xx[asset.category] += 1 
            else:
                xx[asset.category] = 1
        result.append({
            'size':8,
            'plot':bar('数据分布', list(set([i.category for i in assets])), xx)
            })
        return result
