// Felix Almesberger, Pentasys
// | \ O \ |-\

#include "carclient.h"

#include <QCoreApplication>
#include <QEventLoop>
#include <QtCore/QUrl>
#include <QtNetwork/QNetworkRequest>
#include <QtNetwork/QNetworkReply>
#include <QJsonDocument>
#include <QJsonObject>
#include <QException>

CarClient::CarClient()
{
  this->apiEndpoint = "http://localhost:8080/car";
}

void CarClient::startDriving()
{
  QString url = this->apiEndpoint + "/driving/start";
  this->getRestResponse(url);
}

void CarClient::stopDriving()
{
  QString url = this->apiEndpoint + "/driving/stop";
  this->getRestResponse(url);
}

void CarClient::recharge()
{
  QString url = this->apiEndpoint + "/recharge";
  this->getRestResponse(url);
}

CarStatus CarClient::getStatus()
{
  QString url = this->apiEndpoint + "/status";


  QJsonObject response = this->getRestResponse(url);

  CarStatus result = CarStatus();
  result.battery = response["battery_level"].toDouble();
  result.isDriving = response["is_driving"].toBool();
  result.networkReceiption = response["network_reception"].toDouble();
  result.version = response["version"].toDouble();

  return result;
}

QJsonObject CarClient::getRestResponse(QString path)
{
  // QEventLoop will synchronizly wait for finished SIGNAL of the reply
  QEventLoop eventLoop;
  QNetworkAccessManager mgr;
  QObject::connect(&mgr, SIGNAL(finished(QNetworkReply*)), &eventLoop, SLOT(quit()));
  QNetworkRequest request = QNetworkRequest(QUrl(path));
  QNetworkReply *reply = mgr.get(request);
  eventLoop.exec();

  // If something went wrong -> retur empty JSON Object
  if (reply->error() != QNetworkReply::NoError)
  {
    throw QException{};
  }

  // Get Response String
  QString strReply = (QString)reply->readAll();

  // Parse JSON
  QJsonDocument jsonResponse = QJsonDocument::fromJson(strReply.toUtf8());
  return jsonResponse.object();
}

