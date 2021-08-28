# Generated by Django 3.1.7 on 2021-03-14 04:59

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('task', '0004_taskdetail'),
    ]

    operations = [
        migrations.AlterModelOptions(
            name='taskdetail',
            options={'verbose_name': '敏感数据发现详细', 'verbose_name_plural': '敏感数据发现详细'},
        ),
        migrations.AddField(
            model_name='task',
            name='version',
            field=models.IntegerField(default=1, verbose_name='版本'),
        ),
        migrations.AddField(
            model_name='taskdetail',
            name='version',
            field=models.IntegerField(default=1, verbose_name='版本'),
        ),
        migrations.AlterField(
            model_name='task',
            name='create_time',
            field=models.DateTimeField(auto_now=True, verbose_name='创建时间'),
        ),
        migrations.AlterField(
            model_name='task',
            name='update_time',
            field=models.DateTimeField(auto_now_add=True, verbose_name='更新时间'),
        ),
        migrations.AlterField(
            model_name='taskdetail',
            name='create_time',
            field=models.DateTimeField(auto_now=True, verbose_name='创建时间'),
        ),
        migrations.AlterField(
            model_name='taskdetail',
            name='update_time',
            field=models.DateTimeField(auto_now_add=True, verbose_name='更新时间'),
        ),
    ]
