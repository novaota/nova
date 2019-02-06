#ifndef CARCLIENT_H
#define CARCLIENT_H

#include<QJsonObject>

struct CarStatus
{
  double battery;
  double networkReceiption;
  bool isDriving;
  double version;
};

class CarClient
{
public:
  CarClient();
  void startDriving();
  void stopDriving();
  void recharge();
  CarStatus getStatus();
private:
  QJsonObject getRestResponse(QString path);
  QString apiEndpoint;
};

#endif // CARCLIENT_H
