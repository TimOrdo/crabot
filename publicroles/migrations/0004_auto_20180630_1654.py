# Generated by Django 2.0.6 on 2018-06-30 16:54

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('publicroles', '0003_auto_20180630_1652'),
    ]

    operations = [
        migrations.AlterField(
            model_name='publicrole',
            name='uid',
            field=models.CharField(max_length=32, unique=True),
        ),
    ]
