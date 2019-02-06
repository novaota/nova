//import QtQuick.Controls.Universal 2.1
//import QtQuick.Dialogs 1.1
//import Qt.labs.settings 1.0
import QtQuick.Window 2.2
import pentasys.nova.car 1.0
import QtQuick 2.6
import QtQuick.Window 2.2
import QtQuick.Controls 1.6


ApplicationWindow {
    id: window
    width: 800
    height: 480
    visible: true
    //visibility: Window.FullScreen
    title: "NOVA"


    //Settings {
    //    id: settings
    //    property string style: "Universal"
    //}

  //  Pane {
  //      width: 800
  //      height: 480

    ViewModel {
        id: viewModel
    }

    Button {
        id: triggerDriveStopButton
        x: 613
        y: 301
        width: 163
        height: 40
        text: viewModel.isDriving ? qsTr("Anhalten") : qsTr("Fahren")
        checkable: false
        onClicked: viewModel.requestStartStop()
        enabled: viewModel.isConnected
        //Universal.theme: Universal.Dark
    }

    ProgressBar {
        id: networkReceptionProgressBar
        x: 88
        y: 442
        width: 688
        height: 21
        value: viewModel.networkReception
    }

    BusyIndicator {
        id: drivingIndicator
        x: 288
        y: 198
        width: 200
        height: 177
        running: viewModel.isDriving
        visible: true
    }

    Label {
        id: label
        x: 9
        y: 444
        width: 71
        height: 21
        text: qsTr("Mobilfunk:")
        horizontalAlignment: Text.AlignRight
        verticalAlignment: Text.AlignVCenter
    }

    ProgressBar {
        id: batteryProgressBar
        x: 88
        y: 417
        width: 688
        height: 21
        value: viewModel.batteryStatus
    }

    Label {
        id: label1
        x: 9
        y: 417
        width: 71
        height: 21
        text: qsTr("Batteriestatus:")
        horizontalAlignment: Text.AlignRight
        verticalAlignment: Text.AlignVCenter
    }

    Label {
        id: lblVersion
        x: 613
        y: 267
        width: 163
        height: 21
        text: qsTr(viewModel.version)
        wrapMode: Text.WrapAnywhere
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
    }

    Button {
        id: rechargeButton
        x: 613
        y: 350
        width: 163
        height: 40
        text: "Aufladen"
        enabled: viewModel.isConnected
        onClicked: viewModel.recharge()
    }

    Image {
        id: image
        x: 0
        y: 8
        width: 776
        height: 178
        fillMode: Image.PreserveAspectFit
        source: "nova_logo.png"
    }

    Label {
        id: lblIsDriving
        x: 288
        y: 381
        width: 200
        height: 21
        text: viewModel.isConnected
        ? ( viewModel.isDriving
              ? qsTr("Das Auto fährt")
              : qsTr("Das Auto fährt nicht"))
        : qsTr("Nicht verbunden")
        verticalAlignment: Text.AlignVCenter
        horizontalAlignment: Text.AlignHCenter
        wrapMode: Text.WrapAnywhere
    }
}
//}
