# Generated by Django 2.0.6 on 2018-06-30 16:52

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('mydiscord', '0003_auto_20180630_1052'),
    ]

    operations = [
        migrations.AlterField(
            model_name='guild',
            name='name',
            field=models.CharField(max_length=32, unique=True),
        ),
        migrations.AlterField(
            model_name='guild',
            name='uid',
            field=models.CharField(max_length=16, unique=True),
        ),
    ]
