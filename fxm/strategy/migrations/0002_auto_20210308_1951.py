# Generated by Django 3.1.7 on 2021-03-08 11:51

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('strategy', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='strategy',
            name='desc',
            field=models.TextField(blank=True, null=True, verbose_name='描述'),
        ),
        migrations.AlterField(
            model_name='strategy',
            name='name',
            field=models.CharField(max_length=200, unique=True, verbose_name='名称'),
        ),
    ]
