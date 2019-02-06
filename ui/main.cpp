#include <QGuiApplication>
#include <QQmlApplicationEngine>

#include "mainviewmodel.h"

int main(int argc, char *argv[])
{
  QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

  QGuiApplication app(argc, argv);

  //Register ViewModel to QML
  qmlRegisterType<MainViewModel>("pentasys.nova.car", 1, 0, "ViewModel");

  QQmlApplicationEngine engine;
  engine.load(QUrl(QStringLiteral("qrc:/main.qml")));
  if (engine.rootObjects().isEmpty())
    return -1;


  return app.exec();
}
