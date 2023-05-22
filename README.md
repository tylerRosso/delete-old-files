# delete-old-files
You must specify a path and a file age (in days). Examples:<br>
<br>
Delete all files older than 7 days in `C:\Users\User\Document`:
```powershell
.\delete-old-files.exe "C:\Users\User\Documents" 7
```

<br>Delete all files older than 1 day in `C:\Users\User\Downloads`:
```powershell
.\delete-old-files.exe "C:\Users\User\Downloads" 1
```

<br>**Recursive deletes will not be performed.**
