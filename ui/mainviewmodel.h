#ifndef MAINVIEWMODEL_H
#define MAINVIEWMODEL_H

#include <QObject>
#include <QTimer>
#include "carclient.h"

class MainViewModel: public QObject
{
  Q_OBJECT

  Q_PROPERTY(bool isDriving READ isDriving WRITE setIsDriving NOTIFY isDrivingChanged)
  Q_PROPERTY(double batteryStatus READ batteryStatus WRITE setBatteryStatus NOTIFY batteryStatusChanged)
  Q_PROPERTY(double networkReception READ networkReception WRITE setNetworkReception NOTIFY networkReceptionChanged)
  Q_PROPERTY(QString version READ version WRITE setVersion NOTIFY versionChanged)
  Q_PROPERTY(bool isConnected READ isConnected WRITE setIsConnected NOTIFY isConnectedChanged)

public:

  MainViewModel();

  bool isDriving();
  void setIsDriving(bool value);

  double networkReception();
  void setNetworkReception(double value);

  double batteryStatus();
  void setBatteryStatus(double value);

  QString version();
  void setVersion(QString value);

  bool isConnected();
  void setIsConnected(bool value);

signals:
  void isDrivingChanged();
  void networkReceptionChanged();
  void batteryStatusChanged();
  void versionChanged();
  void isConnectedChanged();

public slots:
  void update();
  void recharge();
  void requestStartStop();

private:
  void startPolling();
  void stopPolling();

  void initializeTimer();
  QTimer *timer = NULL;

  bool m_isDriving;
  double m_networkReception;
  double m_batteryStatus;
  QString m_version;
  bool m_isConnected;


  bool updateFromApi;
  CarClient *carClient = NULL;

};

#endif // MAINVIEWMODEL_H

