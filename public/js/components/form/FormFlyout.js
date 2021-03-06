/**
 * Created by languid on 12/26/14.
 */

define([
'kernel',
'react',
'ui/Flyout',
'components/form/helper',
'components/form/FormBody',
'components/form/FormBtns'
],
function (core, React, Flyout, formHelper, FormBody, FormBtns) {

    var FormFlyout = React.createClass({displayName: "FormFlyout",
        render: function () {

            this.flyout = this.props.flyout;

            var titleEl = null;
            if( this.props.title ){
                titleEl = React.createElement("div", {className: "hd"}, this.props.title)
            }
            return (
                React.createElement("div", {className: "mod"}, 
                    titleEl, 
                    React.createElement("div", {className: "bd"}), 
                    React.createElement("div", {className: "ft"})
                )
            )
        }
    });

    return function (formConfig, flyoutConfig, extFlyout) {
        var btns = [{
            text: 'OK',
            className: 'btn-primary',
            click: function () {
                this.submitForm();
            }
        }];

        if (formConfig) {
            if (formConfig.buttons && formConfig.buttons[0] == 'append') {
                formConfig.buttons.splice(0, 1);
                formConfig.buttons = btns.concat(formConfig.buttons)
            } else {
                formConfig.buttons = btns;
            }
        }

        formConfig = $.extend({
            title: '',
            useLabel: false,
            inline: false,
            fields: [],
            buttons: []
        }, formConfig);

        flyoutConfig = $.extend({
            onShow: $.noop,
            onHide: $.noop,
            init: $.noop,
            submit: $.noop,
            classStyle: 'form box'
        }, flyoutConfig);

        extFlyout = $.extend({}, formHelper, extFlyout);

        if (formConfig.inline && formConfig.useLabel) {
            flyoutConfig.classStyle += ' form-horizontal';
        }

        var div = $('<div />', {
            id: 'id'+core.random(10)
        });
        div.data('reactElement', React.render(React.createElement(FormFlyout, {title: formConfig.title}), div[0]));

        var flyout = new Flyout(div, flyoutConfig, extFlyout);

        flyout.body = flyout.element.find('.bd');
        flyout.footer = flyout.element.find('.ft');

        flyout.formBody = React.render(React.createElement(FormBody, {
            fields: formConfig.fields, 
            overload: flyout, 
            inline: formConfig.inline, 
            useLabel: formConfig.useLabel}
        ), flyout.body[0]);

        flyout.formBtns = React.render(React.createElement(FormBtns, {
            buttons: formConfig.buttons, 
            overload: flyout}
        ), flyout.footer[0]);

        return flyout;
    }
});