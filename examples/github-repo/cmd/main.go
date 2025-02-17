package main

import (
	"errors"
	"fmt"

	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn"
	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/cfn/handler"
	"github.com/aws-cloudformation/aws-cloudformation-rpdk-go-plugin/examples/github-repo/cmd/resource"
)

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'cfn generate' command.
*/

type Handler struct{}

func (r *Handler) Create(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Create)
}

func (r *Handler) Read(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Read)
}

func (r *Handler) Update(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Update)
}

func (r *Handler) Delete(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.Delete)
}

func (r *Handler) List(req handler.Request) handler.ProgressEvent {
	return wrap(req, resource.List)
}

// main is the entry point of the applicaton.
func main() {
	cfn.Start(&Handler{})
}

type handlerFunc func(handler.Request, *resource.Model, *resource.Model) (handler.ProgressEvent, error)

func wrap(req handler.Request, f handlerFunc) (response handler.ProgressEvent) {
	defer func() {
		// Catch any panics and return a failed ProgressEvent
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = errors.New(fmt.Sprint(r))
			}

			response = handler.NewFailedEvent(err)
		}
	}()

	// Populate the previous model
	prevModel := &resource.Model{}
	if err := req.UnmarshalPrevious(prevModel); err != nil {
		return handler.NewFailedEvent(err)
	}

	// Populate the current model
	currentModel := &resource.Model{}
	if err := req.Unmarshal(currentModel); err != nil {
		return handler.NewFailedEvent(err)
	}

	response, err := f(req, prevModel, currentModel)
	if err != nil {
		return handler.NewFailedEvent(err)
	}

	return response
}
