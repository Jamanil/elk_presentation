@echo off

:: Запускаем приложение и направляем его вывод в файл, дописывая новые строки
powershell -Command "& {./legacy.exe | Out-File -FilePath log.log -Encoding UTF8 -Append}"