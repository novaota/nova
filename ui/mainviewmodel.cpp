#include "mainviewmodel.h".h"
#include "carclient.h"

#include <QException>
#include<QTimer>

// Constructor

MainViewModel::MainViewModel()
{
  this->carClient = new CarClient();
  this->startPolling();
}

// Properties

double MainViewModel::batteryStatus()
{
  return this->m_batteryStatus;
}

void MainViewModel::setBatteryStatus(double value)
{
  this->m_batteryStatus = value;
  emit batteryStatusChanged();
}

void MainViewModel::requestStartStop()
{
  if(this->m_isDriving)
  {
    this->carClient->stopDriving();
  }
  else
  {
    this->carClient->startDriving();
  }
}

double MainViewModel::networkReception()
{
  return this->m_networkReception;
}

void MainViewModel::setNetworkReception(double value)
{
  this->m_networkReception = value;
  emit networkReceptionChanged();
}

bool MainViewModel::isDriving()
{
  return this->m_isDriving;
}

void MainViewModel::setIsDriving(bool isDriving)
{
  this->m_isDriving = isDriving;
  emit this->isDrivingChanged();
}

QString MainViewModel::version()
{
  return this->m_version;
}

void MainViewModel::setVersion(QString value)
{
  this->m_version = value;
  emit this->versionChanged();
}

bool MainViewModel::isConnected()
{
  return this->m_isConnected;
}

void MainViewModel::setIsConnected(bool value)
{
  this->m_isConnected = value;
  emit this->isConnectedChanged();
}

// Polling Status
void MainViewModel::startPolling()
{
  this->initializeTimer();
  this->timer->start();
}

void MainViewModel::initializeTimer()
{
  this->setIsConnected(false);
  this->timer = new QTimer(this);
  this->timer->setInterval(1000);
  connect(this->timer, SIGNAL(timeout()), this, SLOT(update()));
}

void MainViewModel::update()
{
  this->timer->stop();
  try
  {
    CarStatus status = this->carClient->getStatus();
    this->updateFromApi = true;
    this->setBatteryStatus(status.battery);
    this->setIsDriving(status.isDriving);
    this->setNetworkReception(status.networkReceiption);
    this->setVersion(QString::number(status.version));
    this->setIsConnected(true);

    this->updateFromApi = false;
  }
  catch(QException &e)
  {
    this->setIsConnected(false);
    this->setIsDriving(false);
  }

  this->updateFromApi = false;
  this->timer->start();
}

void MainViewModel::recharge()
{
  this->carClient->recharge();
}

void MainViewModel::stopPolling()
{
  this->timer->stop();
  delete this->timer;
  this->timer = NULL;
}
