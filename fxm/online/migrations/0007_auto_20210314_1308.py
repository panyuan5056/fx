# Generated by Django 3.1.7 on 2021-03-14 05:08

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('online', '0006_auto_20210308_1951'),
    ]

    operations = [
        migrations.AlterField(
            model_name='online',
            name='server',
            field=models.CharField(max_length=200, verbose_name='探针地址'),
        ),
        migrations.AlterField(
            model_name='online',
            name='token',
            field=models.TextField(blank=True, max_length=200, null=True, verbose_name='探针token'),
        ),
    ]