@echo off

::��.txt�ļ�ת��Ϊ��Ӧ��.go�ļ�

::��ȡ.txt�ļ���·�� ���ȡ��Ŀ¼�������ļ����µ�.txt�ļ�
set readPath=F:\code\src\makeMsg\test
::����.go�ļ���·�� ������Ϸ���ȡ·�����ļ����浽��Ӧ�ļ�����
set writePath=F:\code\src\makeMsg\tool
::���ɿ�ִ���ļ���·��
set exePath=F:\code\src\makeMsg

echo ��ʼ����.go�ļ�
%exePath%\makeMsg -writePath=%writePath% -readPath=%readPath%

echo ������ϣ������������
TIMEOUT /T 999